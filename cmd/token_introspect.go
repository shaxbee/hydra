/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var tokenIntrospectCmd = &cobra.Command{
	Use:   "introspect <token>",
	Short: "Introspect an access or refresh token",
	Run:   cmdHandler.Introspection.Introspect,
}

func init() {
	tokenCmd.AddCommand(tokenIntrospectCmd)
	tokenIntrospectCmd.Flags().StringSlice("scope", []string{}, "Additionally check if scope was granted")
	tokenIntrospectCmd.Flags().String("endpoint", os.Getenv("HYDRA_ADMIN_URL"), "Set the URL where ORY Hydra is hosted, defaults to environment variable HYDRA_ADMIN_URL")
	tokenIntrospectCmd.Flags().String("client-id", os.Getenv("OAUTH2_CLIENT_ID"), "Use the provided OAuth 2.0 Client ID, defaults to environment variable OAUTH2_CLIENT_ID")
	tokenIntrospectCmd.Flags().String("client-secret", os.Getenv("OAUTH2_CLIENT_SECRET"), "Use the provided OAuth 2.0 Client Secret, defaults to environment variable OAUTH2_CLIENT_SECRET")
	tokenIntrospectCmd.Flags().String("access-token", os.Getenv("OAUTH2_ACCESS_TOKEN"), "Set an access token to be used in the Authorization header, defaults to environment variable OAUTH2_ACCESS_TOKEN")
}
