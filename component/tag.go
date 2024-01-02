package component

import (
	"context"
	"fmt"
	"log/slog"

	"opencsg.com/starhub-server/builder/store/database"
	"opencsg.com/starhub-server/common/config"
	"opencsg.com/starhub-server/component/tagparser"
)

func NewTagComponent(config *config.Config) (*TagComponent, error) {
	tc := &TagComponent{}
	tc.ts = database.NewTagStore()
	return tc, nil
}

type TagComponent struct {
	ts *database.TagStore
}

func (tc *TagComponent) AllTags(ctx context.Context) ([]database.Tag, error) {
	//TODO: query cache for tags at first
	return tc.ts.AllTags(ctx)
}

func (c *TagComponent) ClearMetaTags(ctx context.Context, namespace, name string) error {
	_, err := c.ts.SetMetaTags(ctx, namespace, name, nil)
	return err
}

func (c *TagComponent) UpdateMetaTags(ctx context.Context, tagScope database.TagScope, namespace, name, content string) ([]*database.RepositoryTag, error) {

	var tp tagparser.TagProcessor
	//TODO:load from cache
	if tagScope == database.DatasetTagScope {
		tp = tagparser.NewDatasetTagProcessor(c.ts)
	} else {
		tp = tagparser.NewModelTagProcessor(c.ts)
	}
	tagsMatched, tagToCreate, err := tp.ProcessReadme(ctx, content)
	if err != nil {
		slog.Error("Failed to process tags", slog.Any("error", err))
		return nil, fmt.Errorf("failed to process tags, cause: %w", err)
	}
	slog.Debug("tagsToCreate", slog.Any("tags", tagToCreate))
	slog.Debug("tagsMatched", slog.Any("tags", tagsMatched))

	err = c.ts.SaveTags(ctx, tagToCreate)
	if err != nil {
		slog.Error("Failed to save tags", slog.Any("error", err))
		return nil, fmt.Errorf("failed to save tags, cause: %w", err)
	}
	metaTags := append(tagsMatched, tagToCreate...)
	var repoTags []*database.RepositoryTag
	repoTags, err = c.ts.SetMetaTags(ctx, namespace, name, metaTags)
	if err != nil {
		slog.Error("failed to set dataset's tags", slog.String("namespace", namespace),
			slog.String("name", name), slog.Any("error", err))
		return nil, fmt.Errorf("failed to set dataset's tags, cause: %w", err)
	}

	return repoTags, nil
}

func (c *TagComponent) UpdateLibraryTags(ctx context.Context, tagScope database.TagScope, namespace, name, oldFilePath, newFilePath string) error {
	oldLibTagName := tagparser.LibraryTag(oldFilePath)
	newLibTagName := tagparser.LibraryTag(newFilePath)
	//TODO:load from cache
	var (
		allTags []*database.Tag
		err     error
	)
	if tagScope == database.DatasetTagScope {
		allTags, err = c.ts.AllDatasetTags(ctx)
	} else {
		allTags, err = c.ts.AllModelTags(ctx)
	}
	if err != nil {
		return fmt.Errorf("failed to get all tags, error: %w", err)
	}
	var oldLibTag, newLibTag *database.Tag
	for _, t := range allTags {
		if t.Name == newLibTagName {
			newLibTag = t
		}
		if t.Name == oldLibTagName {
			oldLibTag = t
		}
	}
	err = c.ts.SetLibraryTag(ctx, namespace, name, newLibTag, oldLibTag)
	if err != nil {
		slog.Error("failed to set dataset's tags", slog.String("namespace", namespace),
			slog.String("name", name), slog.Any("error", err))
		return fmt.Errorf("failed to set Library tags, cause: %w", err)
	}
	return nil
}
