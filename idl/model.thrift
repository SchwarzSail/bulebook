namespace go model

struct User {
    1: string account //账号
    2: string username
    3: string major //专业
    4: string email //邮箱
    5: string avatar //头像
    6: string role
}

struct Project {
    1: string title
    2: string description
    3: string username //发起人
    4: string numbers //参与人数
    5: string types//专业类型
}