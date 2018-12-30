package controllers

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {}
func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}
func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}