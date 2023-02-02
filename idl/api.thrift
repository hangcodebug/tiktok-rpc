namespace go api

struct BaseResp{
    1:i64 status_code
    2:string status_msg
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
  1:i64 user_id  (api.query="user_id", api.vd="len($) > 0")
  2:string token (api.query="token", api.vd="len($) > 0")
}

struct douyin_user_response{
    1:BaseResp base_resp
    2:User user
}
// 注册用户
struct douyin_user_register_request{
    1:string username (api.query="username", api.vd="len($) > 0")
    2:string password (api.query="password", api.vd="len($) > 0")
}
struct douyin_user_register_response{
    1:BaseResp base_resp
    2:i64 user_id // 用户id
    3:string token // 用户鉴权token
}

service ApiService{
    douyin_user_response GetUserInfo(1:douyin_user_request req) (api.get="douyin/user")
    douyin_user_register_response RegisterUser(1:douyin_user_register_request req) (api.post="douyin/user/register")
}
