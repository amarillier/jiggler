package main

import (
	updatechecker "github.com/amarillier/go-update-checker"
)

func updateChecker(repoOwner string, repo string, repoName string, repodl string) (string, bool) {
	// uc := updatechecker.New("amarillier", "jiggler", "jiggler", "", 1, false)
	uc := updatechecker.New(repoOwner, repo, repoName, repodl, 0, false)
	uc.CheckForUpdate(appVersion)
	// uc.PrintMessage()
	updtmsg := uc.Message
	return updtmsg, uc.UpdateAvailable
}

// "Now this is not the end. It is not even the beginning of the end. But it is, perhaps, the end of the beginning." Winston Churchill, November 10, 1942
