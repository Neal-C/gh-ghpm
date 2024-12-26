package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

// STARS_THRESHOLD : the required numbers of stars on a repository for it be avoided by ghpm
const STARS_THRESHOLD uint = 1

type User struct {
	Username string `json:"login"`
}

type GithubRepository struct {
	Stars uint `json:"stargazers_count"`

	Fullname string `json:"full_name"`
}

func Prettyfy(data any) (string, error) {
	val, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

var thanosSnapCmd = &cobra.Command{
	Use:   "thanos_snap",
	Short: "Switch all your public repositories to private.",
	Args:  cobra.NoArgs,
	Long: heredoc.Docf(`
		Switch all your public repositories to private.

		By default, starred repositories with 1 stars are not turned private.
	`, "`"),
	Example: heredoc.Doc(`
		# request all your public repositories to turn private, except if they have stars
		
		$ ghpm thanos_snap
		`),
	RunE: func(cmd *cobra.Command, args []string) error {

		client, err := api.DefaultRESTClient()
		if err != nil {
			return err
		}

		var user User

		err = client.Get("user", &user)

		if err != nil {
			return err
		}

		fmt.Printf("running as %s\n", user.Username)

		publicRepositoriesGithubAPIEndpoint := fmt.Sprintf("https://api.github.com/users/%s/repos?visibility=public&per_page=100", user.Username)

		readmeRepository := fmt.Sprintf("%s/%s", user.Username, user.Username)

		payload := map[string]any{
			"private": true,
		}

		jsonPayload, err := json.Marshal(payload)

		if err != nil {
			return fmt.Errorf("json.Marshal: %s", err)
		}

		for {

			httpResponse, err := client.Request(http.MethodGet, publicRepositoriesGithubAPIEndpoint, nil)

			if err != nil {
				return fmt.Errorf("something with the Github API went wrong: %s", err)
			}

			var publicRepositories []GithubRepository

			if err := json.NewDecoder(httpResponse.Body).Decode(&publicRepositories); err != nil {
				return err
			}

			httpResponse.Body.Close()

			var namesOfPublicRepositories = make([]string, 0, 100)

			for _, repo := range publicRepositories {
				namesOfPublicRepositories = append(namesOfPublicRepositories, repo.Fullname)
			}

			names, err := Prettyfy(namesOfPublicRepositories)

			if err != nil {
				return err
			}

			fmt.Printf("your public repositories : %s \n", names)

			// TODO : lobby github for a batch request endpoint, so that it can be only 1 HTTP call and not O(n) HTTP calls
			for _, repo := range publicRepositories {

				if repo.Stars >= STARS_THRESHOLD {

					log.Printf("repository %s cannot be switched to private by ghpm because it has more than %d stars (%d exactly) \n", repo.Fullname, STARS_THRESHOLD, repo.Stars)

					continue
				}

				if repo.Fullname == readmeRepository {

					fmt.Printf("dodging the README repository %s because it's a special repository \n", readmeRepository)

					continue
				}

				currentPublicRepositoryEndpoint := fmt.Sprintf("https://api.github.com/repos/%s", repo.Fullname)

				httpResponse, err := client.RequestWithContext(cmd.Context(), http.MethodPatch, currentPublicRepositoryEndpoint, bytes.NewBuffer(jsonPayload))

				if err != nil {

					log.Printf("error requesting %s: %s \n", repo.Fullname, err)
					log.Println("skipping", repo.Fullname)

					continue
				}

				switch {
				case httpResponse.StatusCode == http.StatusNotImplemented:

					log.Printf("%s was not switched to private. I suggest to you try from the web version for this one. I am sorry for failing you, please complain to the developer \n", repo.Fullname)

					httpResponse.Body.Close()

					continue

				case httpResponse.StatusCode == http.StatusNotFound:

					log.Printf("%s was not switched to private. Because it was not found. Did you misspell?\n", repo.Fullname)

					httpResponse.Body.Close()

					continue

				case httpResponse.StatusCode >= 500:

					log.Printf("github is likely down. Retry. If it does persist: Please complain to the developer \n")

					httpResponse.Body.Close()

					continue
				}

				log.Printf("%s switched to private. \n", repo.Fullname)

				httpResponse.Body.Close()

			}

			if len(publicRepositories) != 100 {
				break
			}

		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(thanosSnapCmd)
}
