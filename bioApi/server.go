package main

import (
	"goBio/config" //menggunakan package config untuk menghubungkan ke database
	"goBio/controllers"
	"goBio/model" //menggunakan package model untuk mengambil data dari database

	//menggunakan package model untuk membuat, mengubah, dan menghapus data di database
	"goBio/utils"   //menggunakan package utils untuk membuat nama file baru secara random
	"io"            //menggunakan package io untuk menyalin isi file yang diupload ke file baru
	"os"            //menggunakan package os untuk membuat file baru
	"path/filepath" //menggunakan package filepath untuk mengambil ekstensi file yang diupload
	"strings"       //menggunakan package strings untuk mengambil header Content-Type

	"github.com/labstack/echo/v4" //menggunakan echo versi 4
)

func main() {
	config.ConnectDB()  //menghubungkan ke database
	route := echo.New() //membuat variabel route untuk menampung echo.New() / echo new untuk membuat instance baru dari echo

	// handle open image from browser with path /uploads
	route.Static("/uploads", "uploads") //digunakan untuk mengakses file yang ada di folder uploads

	// route untuk menampilkan hello world
	route.GET("/", func(c echo.Context) error {
		data := "Hello World REST API with Golang"
		return c.JSON(200, data) //mengembalikan response dengan status code 200 dan data berupa string
	})

	// route untuk membuat user baru
	route.POST("user/create_user", func(c echo.Context) error {
		user := new(model.Users) //digunakan untuk membuat variabel user
		c.Bind(user)
		contentType := c.Request().Header.Get("Content-Type") //digunakan untuk mengambil header Content-Type
		response := new(model.Response)                       //digunakan untuk membuat variabel response

		if user.Email == "" || user.Nama == "" || user.NoHp == "" || user.Alamat == "" { //jika email, nama, no_hp, dan alamat kosong, maka akan mengembalikan response error
			response.Status = false
			response.Message = "Gagal menyimpan data, email, nama, no_hp, dan alamat harus diisi"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else if _, err := controllers.GetOneByEmail(user.Email); err == nil { //jika email sudah terdaftar, maka akan mengembalikan response error
			response.Status = false
			response.Message = "Gagal menyimpan data, email sudah terdaftar"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else if strings.Contains(contentType, "multipart/form-data") ||
			strings.Contains(contentType, "application/x-www-form-urlencoded") { //jika header Content-Type adalah multipart/form-data atau application/x-www-form-urlencoded, maka akan mengembalikan response error
			file, err := c.FormFile("ktp") //digunakan untuk mengambil file yang diupload
			if err != nil {
				response.Status = false
				response.Message = "Gagal menyimpan data, file KTP tidak ditemukan"
				response.ErrorCode = 400
				return c.JSON(400, response)
			} else {
				src, err := file.Open() //digunakan untuk membuka file yang diupload
				if err != nil {
					response.Status = false
					response.Message = "Gagal menyimpan data, file KTP error saat dibuka"
					response.ErrorCode = 400
					return c.JSON(400, response)
				}
				defer src.Close() //digunakan untuk menutup file yang dibuka

				fileExt := filepath.Ext(file.Filename)   //digunakan untuk mengambil ekstensi file yang diupload
				randomFileName := utils.RandomString(10) //digunakan untuk membuat nama file baru secara random
				// check jika file name sudah ada, maka buat file name baru
				if _, err := os.Stat("uploads/" + randomFileName + fileExt); err == nil {
					randomFileName = utils.RandomString(10)
				}

				newFileName := randomFileName + fileExt //digunakan untuk menggabungkan nama file baru dengan ekstensi file yang diupload

				dst, err := os.Create("uploads/" + newFileName) //digunakan untuk membuat file baru
				if err != nil {
					response.Status = false
					response.Message = "Gagal menyimpan data, file KTP error saat dibuat"
					response.ErrorCode = 400
					return c.JSON(400, response)
				}
				defer dst.Close() //digunakan untuk menutup file yang dibuat

				if _, err = io.Copy(dst, src); err != nil { //digunakan untuk menyalin isi file yang diupload ke file baru
					response.Status = false
					response.Message = "Gagal menyimpan data, file KTP error saat disalin"
					response.ErrorCode = 400
					return c.JSON(400, response)
				}

				user.Ktp = newFileName //digunakan untuk menyimpan nama file baru ke variabel user.Ktp
			}
		}

		// create user if all conditions are met
		if err := controllers.CreateUser(user); err != nil {
			response.Status = false
			response.Message = "Gagal menyimpan data, terjadi kesalahan dikarenakan " + err.Error()
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else {
			response.Status = true
			response.Message = "Berhasil menyimpan data"
			response.Data = user
			response.ErrorCode = 200
			return c.JSON(200, response)
		}

	})

	//route untuk mengambil user berdasarkan email
	route.PUT("user/update_user/:email", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		contentType := c.Request().Header.Get("Content-Type") //digunakan untuk mengambil header Content-Type
		response := new(model.Response)
		if user.Email == "" || user.Nama == "" || user.NoHp == "" || user.Alamat == "" {
			response.Status = false
			response.Message = "Gagal menyimpan data, email, nama, no_hp, dan alamat harus diisi"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else if !strings.Contains(contentType, "application/json") { // response error if header Content-Type is application/json
			response.Status = false
			response.Message = "Gagal menyimpan data, header Content-Type harus berupa application/json"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else if user.Email != c.Param("email") {
			response.Status = false
			response.Message = "Gagal menyimpan data, email tidak boleh diubah"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else if _, err := controllers.GetOneByEmail(user.Email); err != nil {
			response.Status = false
			response.Message = "Gagal menyimpan data, email tidak ditemukan"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else {
			//jangan simpan user.Ktp jika header Content-Type adalah application/json
			//buat variabel tempUser untuk menampung data user yang sudah ada di database
			//check if error is nil, if not nil then return error
			tempUser, err := controllers.GetOneByEmail(user.Email)
			if err != nil {
				response.Status = false
				response.Message = "Gagal menyimpan data, email tidak ditemukan"
				response.ErrorCode = 400
				return c.JSON(400, response)
			}

			// check user.Ktp, berbeda dengan di database, maka user.Ktp tidak boleh diubah
			if (user.Ktp != "" && tempUser.Ktp != "") && user.Ktp != tempUser.Ktp {
				response.Status = false
				response.Message = "Gagal menyimpan data, file KTP tidak boleh diubah"
				response.ErrorCode = 400
				return c.JSON(400, response)
			}

			// update user.Ktp if header Content-Type is multipart/form-data or application/x-www-form-urlencoded
			user.Ktp = tempUser.Ktp
			user.CreatedAt = tempUser.CreatedAt

			// update user if all conditions are met
			if err := controllers.UpdateUser(user, c.Param("email")); err != nil {
				response.Status = false
				response.Message = "Gagal menyimpan data, terjadi kesalahan dikarenakan " + err.Error()
				response.ErrorCode = 400
				return c.JSON(400, response)
			} else {
				response.Status = true
				response.Message = "Berhasil menyimpan data"
				response.Data = user
				response.ErrorCode = 200
				return c.JSON(200, response)
			}
		}
	})

	//route untuk menghapus user berdasarkan email
	route.DELETE("user/delete_user/:email", func(c echo.Context) error {
		user, _ := controllers.GetOneByEmail(c.Param("email"))
		response := new(model.Response)

		if user.Email == "" {
			response.Status = false
			response.Message = "Gagal menghapus data, email tidak ditemukan"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else {
			if err := controllers.DeleteUser(&user); err != nil {
				response.Status = false
				response.Message = "Gagal menghapus data, email tidak ditemukan"
				response.ErrorCode = 400
				return c.JSON(400, response)
			} else {
				// check if file exists, after delete file ktp
				if _, err := os.Stat("uploads/" + user.Ktp); err == nil {
					if err := os.Remove("uploads/" + user.Ktp); err != nil {
						response.Status = false
						response.Message = "Gagal menghapus data, file KTP error saat dihapus"
						response.ErrorCode = 400
						return c.JSON(400, response)
					}
				}

				response.Status = true
				response.Message = "Berhasil menghapus data, dengan email " + user.Email
				response.ErrorCode = 200
				return c.JSON(200, response)
			}
		}
	})

	//route untuk mengambil user berdasarkan email
	route.GET("user/search_user", func(c echo.Context) error {
		response := new(model.Response)
		users, err := controllers.GetAll(c.QueryParam("keyword"))
		if err != nil {
			response.Status = false
			response.Message = "Gagal mengambil data, keyword tidak ditemukan"
			response.ErrorCode = 400
			return c.JSON(400, response)
		} else {
			// customr response path for ktp
			for i := 0; i < len(users); i++ {
				users[i].Ktp = "/uploads/" + users[i].Ktp
			}

			response.Status = true
			response.Message = "Berhasil mengambil data"
			response.Data = users
			response.ErrorCode = 200
			return c.JSON(200, response)
		}
	})

	//route start pada port 9000
	route.Start(":9000")

}
