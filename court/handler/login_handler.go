package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/Surafeljava/Court-Case-Management-System/entity"
)

type LoginHandler struct {
	tmpl     *template.Template
	loginSrv caseUse.LoginService
}

func NewLoginHandler(T *template.Template, LS caseUse.LoginService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginSrv: LS}
}

func (lh *LoginHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		user_id := r.FormValue("user_id")
		user_pwd := r.FormValue("user_pwd")

		fmt.Println(user_id)
		fmt.Println(user_pwd)

		who := CheckWho(user_id)

		error_message := entity.SuccessMessage{Status: "Error", Message: "Wrong ID or Password Try again"}
		success_message := entity.SuccessMessage{Status: "Success", Message: "Login Success!"}

		if who == 0 {
			adm, err := lh.loginSrv.CheckAdmin(user_id, user_pwd)
			if adm != nil {
				// togo := struct {
				// 	admin entity.Admin
				// 	msg   entity.SuccessMessage
				// }{
				// 	*adm,
				// 	success_message,
				// }

				lh.tmpl.ExecuteTemplate(w, "admin.home.layout", adm)
			} else if len(err) > 0 {
				lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
			}
		} else if who == 1 {
			jud, err := lh.loginSrv.CheckJudge(user_id, user_pwd)
			if jud != nil {
				togo := struct {
					judge entity.Judge
					msg   entity.SuccessMessage
				}{
					*jud,
					success_message,
				}
				lh.tmpl.ExecuteTemplate(w, "judge.home.layout", togo)
			} else if len(err) > 0 {
				lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
			}
		} else if who == 2 {
			opp, err := lh.loginSrv.CheckOpponent(user_id, user_pwd)
			if opp != nil {
				togo := struct {
					opponent entity.Opponent
					msg      entity.SuccessMessage
				}{
					*opp,
					success_message,
				}
				lh.tmpl.ExecuteTemplate(w, "opponent.home.layout", togo)
			} else if len(err) > 0 {
				lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
			}
		} else {
			lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
		}

	} else {
		lh.tmpl.ExecuteTemplate(w, "login.layout", nil)
	}

}

func CheckWho(id string) int {
	check := id[0:2]
	fmt.Println(check)
	if check == "AD" {
		return 0
	} else if check == "JU" {
		return 1
	} else if check == "OP" {
		return 2
	}
	return -1
}
