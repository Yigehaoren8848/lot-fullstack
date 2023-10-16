/**
 * 配置编译环境和线上环境之间的切换
 * 
 * baseUrl: 域名地址
 * routerMode: 路由模式
 * imgBaseUrl: 图片所在域名地址
 * 
 */

let baseUrl = ''; 
let routerMode = 'hash';
let imgBaseUrl = '';


if (process.env.NODE_ENV == 'development') {
    imgBaseUrl = '/img/';
	baseUrl = 'http://localhost:8080'

}else if(process.env.NODE_ENV == 'production'){
	baseUrl = 'http://localhost:8080'
    imgBaseUrl = '';
}

export {
	baseUrl,
	routerMode,
	imgBaseUrl,
}