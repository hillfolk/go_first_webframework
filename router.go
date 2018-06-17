package main

import "net/http"
import "strings"

type router struct {
	//키: http method
	//값: URL 에 대응되는 handler func
	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지확인
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	m[pattern] = h
}

func (r *router) handler() HandlerFunc {
	return func(c *Context) {
		for pattern, handler := range r.handlers[c.Request.Method] {
			if ok, params := match(pattern, c.Request.URL.Path); ok {
				for k, v := range params {
					c.Params[k] = v
				}
				handler(c)
				return
			}
		}
		http.NotFound(c.ResponseWriter, c.Request)
		return
	}
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	// 패턴과 패스르 / 단위로 분리
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	// 패턴과 패스의 항목 수가 맞지 않으면 바로 false 반환
	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	// "/"로 구분된 패턴/ 패스의 각 문자열을 하나씩 비교

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			// 문자열이 일치하면 루프 실행
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			// 패턴이 ":"로 시작하면 params 에 URL params 를 담은후 루프 수행
			params[patterns[i][1:]] = paths[i]

		default:
			// 일치하는 경우가 없으면 False 반환
			return false, nil
		}
	}

	//True 와 params 를 반환
	return true, params

}

