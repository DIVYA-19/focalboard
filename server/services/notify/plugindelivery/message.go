// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package plugindelivery

import (
	"fmt"

	"github.com/mattermost/focalboard/server/model"
)

const (
	// TODO: localize these when i18n is available.
	defCommentTemplate     = "@%s mentioned you in a comment on the card [%s](%s)\n> %s"
	defDescriptionTemplate = "@%s mentioned you in the card [%s](%s)\n> %s"
)

func formatMessage(author string, extract string, card string, link string, block *model.Block) string {
	template := defDescriptionTemplate
	if block.Type == model.TypeComment {
		template = defCommentTemplate
	}
	return fmt.Sprintf(template, author, card, link, extract)
}

func makeLink(serverRoot string, workspace string, board string, card string) string {
	return fmt.Sprintf("%s/workspace/%s/%s/0/%s/", serverRoot, workspace, board, card)
}
