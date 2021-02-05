import request from '@/utils/request'

export function fetchList(params) {
  
    return request({
        url: '/articles',
        method: 'get',
        params: params
    })
}

export function fetchArticle(params) {

  return request({
    url: '/article',
    method: 'get',
    params: params
  })
}

export function fetchFocus() {
    return request({
        url: '/top_articles',
        method: 'get',
        params: {}
    })
}

export function fetchCategory() {
    return request({
        url: '/all_tags',
        method: 'get',
        params: {}
    })
}

export function fetchFriend() {
    return request({
        url: '/friend',
        method: 'get',
        params: {}
    })
}

export function fetchSocial() {
    return request({
        url: '/social',
        method: 'get',
        params: {}
    });
}

export function fetchSiteInfo() {
    return request({
        url: '/site_config',
        method: 'get',
        params: {}
    })
}

export function fetchComment() {
    return request({
        url: '/comment',
        method: 'get',
        params: {}
    })
}
