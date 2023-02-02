namespace go userCore

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResponse {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct User{
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}
// 获得用户信息
struct douyin_user_request{
  1:i64 user_id
  2:string token
}

struct douyin_user_response{
    1:BaseResponse base_resp
    2:User user
}
// 注册用户
struct douyin_user_register_request{
    1:string username (vt.min_size = "1")
    2:string password (vt.min_size = "1")
}
struct douyin_user_register_response{
    1:BaseResponse base_resp
    2:i64 user_id // 用户id
    3:string token // 用户鉴权token
}

service UserCoreService{
    douyin_user_response GetUserInfo(1:douyin_user_request req)
    douyin_user_register_response RegisterUser(1:douyin_user_register_request req)
}