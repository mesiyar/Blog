import request from '../utils/request';

export const fetchData = query => {
  return request({
    url: './table.json',
    method: 'get',
    params: query
  });
};

export const fetchTags = query => {
  return request({
    url: '/admin/tags',
    method: 'get',
    params: query
  });
};


export const fetchArticle = query => {
  return request({
    url: '/admin/articles',
    method: 'get',
    params: query
  });
};

export const fetchConfigs = query => {
  return request({
    url: '/admin/configs',
    method: 'get',
    params: query
  });
};