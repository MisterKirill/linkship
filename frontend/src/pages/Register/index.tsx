import { FormEvent, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { authenticate } from "../../utils/users"
import classes from "./style.module.css"

function Register() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const navigate = useNavigate()

  const onSubmit = async (event: FormEvent) => {
    event.preventDefault()

    if (await authenticate({ username, password }, true)) {
      navigate("/")
    }
  }

  return (
    <div className={classes.wrapper}>
      <div>
        <h1>Добро пожаловать на <span className="text-brand">Linkship</span>!</h1>
        <span>Уже есть аккаунт? <Link to="/login" className="link">Войти</Link>.</span>
      </div>

      <form className="form" onSubmit={onSubmit}>
        <input
          type="text"
          placeholder="Имя пользователя"
          className="input"
          onChange={(e) => setUsername(e.target.value)}
          required />
        <input
          type="password"
          placeholder="Пароль"
          className="input"
          onChange={(e) => setPassword(e.target.value)}
          required />
        <button type="submit" className="btn">
          Создать аккаунт
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
            <path fillRule="evenodd" d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708"/>
          </svg>
        </button>
      </form>
    </div>
  )
}

export default Register
