package auth

type AuthLevel string

const AdminAuth AuthLevel = "AdminAuth"
const UserAuth AuthLevel = "UserAuth"
const NoAuth AuthLevel = "NoAuth"
