package auth

type AuthLevel string

const SuperAuth AuthLevel = "SuperAuth"
const AdminAuth AuthLevel = "AdminAuth"
const UserAuth AuthLevel = "UserAuth"
const NoAuth AuthLevel = "NoAuth"