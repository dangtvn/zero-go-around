package posts

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/api/internal/logic/posts"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
	"net/http"
)

func GetPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := posts.NewGetPostsLogic(r.Context(), svcCtx)
		resp, err := l.GetPosts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
