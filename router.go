package main

import "net/http"

type router struct {
	//키: http method
	//값: URL 에 대응되는 handler func
	handlers map[string]map[string]http.HandlerFunc
}


func (* router) HandlerFunc(method, pattern string, h http.HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지확인
	m, ok = r.handlers[method]
	if !ok {
		// 등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}
	m[pattern] = h
}


type Hander interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	
}

func (r *router) ServeHTTP(w http.ResponseWrite, req *http.Request) {
	if m, ok := r.handlers[req.Method]; ok {
		if h, ok := m[req.URL,PATH]; ok {
			// 요청 URL 핸들러 수행
			h(w,req)
			return
		}
	}
	http.NotFound(w, req)
}

func match(pattern, path string) (boo, map[string]string) {
	if pattern == path {
		return true,nil
	}

	patterns := strings.Split(pattern,"/")
	paths := strings.Split(path,"/")

	if len(patterns) != len(paths){
		return flase,nil
	}
}
