package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/post"
	"strconv"
	"time"
)

type PostBaseInterop struct {
	postUseCase post.PostUseCase
	authUseCase auth.AuthUseCase
}

func (p PostBaseInterop) List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	return p.postUseCase.List(ctx, opts)
}

func (p PostBaseInterop) Create(ctx context.Context, token string, data *post.Post) error {
	record, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return err
	}
	data.CreatorId = record.UID
	data.ID = record.UID[:10] + strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	data.Comment = make([]string, 0)
	data.Like = make([]string, 0)
	if data.HashTag == nil {
		data.HashTag = make([]string, 0)
	}
	if data.Mention == nil {
		data.Mention = make([]string, 0)
	}
	data.Share = make([]string, 0)
	data.Status = "active"
	return p.postUseCase.Create(ctx, data)
}

func (p PostBaseInterop) Delete(ctx context.Context, token string, id string) error {
	_, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return err
	}
	return p.postUseCase.Delete(ctx, id)

}

func (p PostBaseInterop) GetDetail(ctx context.Context, token string, id string) (*post.Post, error) {
	_, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return p.postUseCase.GetDetail(ctx, id)
}

func (p PostBaseInterop) GetByUid(ctx context.Context, token string, opts *common.QueryOpts, style string) (*common.ListResult[*post.Post], error) {
	record, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}

	return p.postUseCase.GetByUid(ctx, record.UID, opts, style)
}

func (p PostBaseInterop) GetOther(ctx context.Context, token string, uid string, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	_, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return p.postUseCase.GetOther(ctx, uid, opts)

}
func (p PostBaseInterop) GetByCategory(ctx context.Context, token string, categoryId string, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	_, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return p.postUseCase.GetByCategory(ctx, categoryId, opts)

}

func (p PostBaseInterop) Update(ctx context.Context, token string, data *post.Post) error {
	_, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return err
	}
	postData, err := p.postUseCase.GetDetail(ctx, data.ID)
	if postData.ID != data.ID {
		return post.ErrPostRequiredID
	}
	if postData.CreatorId != data.CreatorId {
		return post.ErrPostRequiredCreatorID
	}
	if data.Comment == nil {
		data.Comment = postData.Comment
	}
	if data.Like == nil {
		data.Like = postData.Like
	}
	if data.HashTag == nil {
		data.HashTag = postData.HashTag
	}
	if data.Mention == nil {
		data.Mention = postData.Mention
	}
	if data.Share == nil {
		data.Share = postData.Share
	}
	data.Status = "active"
	return p.postUseCase.Update(ctx, data)
}

func (p PostBaseInterop) UpdatePostComment(ctx context.Context, token string, id string, data *post.Post) error {
	_, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return err
	}
	postCommentData, err := p.postUseCase.GetDetail(ctx, id)
	if err != nil {
		return err
	}
	postCommentData.ID = id
	postCommentData.Comment = append(postCommentData.Comment, data.Comment...)
	return p.postUseCase.UpdatePostComment(ctx, id, data)
}

func NewPostBaseInterop(postUseCase post.PostUseCase, authUseCase auth.AuthUseCase) PostBaseInterop {
	return PostBaseInterop{
		postUseCase: postUseCase,
		authUseCase: authUseCase,
	}
}
