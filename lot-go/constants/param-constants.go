package constants

const JWTSECREETKEY = "JWTForKey9982"

// 处理正常
const SuccessCode = 1

// 参数异常
const ParamErrorCode = 20001

// 数据库不存在某一条数据
const ModelNotExistCode = 10002

// 数据库已存在某一条数据
const ModelHasExistCode = 10002

// 密码错误
const PasswordErrorCode = 10003

// 系统异常，token生成失败
const TokenGenerateError = 10003

// 系统异常，token生成失败
const InsertUserError = 10004

// 系统异常,获取请求体参数失败
const GainPostDataError = 10005

// 系统异常,数据添加失败
const InsertError = 10006

// 系统异常,查询异常
const QueryError = 10007

// 权限异常
const AuthError = 30001
