package internal

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BasePath = userHomeDir + "/byeoru"

// github username and repo in order to fetch snippets.
var Owner = "hoehwa"
var Repository = "audit"

// urls for hosted websites.
const Baseurl = "https://garden.mindulle.vercel.app/"
