package constants

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/byeoru"

// github username and repo in order to fetch snippets.
var Owner = "hoehwa"
var Repository = "audit"

// urls for hosted websites.
const GardenBaseurl = "https://garden.mindulle.vercel.app/"
