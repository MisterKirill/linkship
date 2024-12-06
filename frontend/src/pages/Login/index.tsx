import { Link } from 'react-router-dom'
import classes from './style.module.css'

function Login() {
  return (
    <div className={classes.wrapper}>
      <div>
        <h1>Добро пожаловать на <span className="text-brand">Linkship</span>!</h1>
        <span>Ещё нет аккаунта? <Link to="/register" className="link">Создать аккаунт</Link>.</span>
      </div>

      <form className="form">
        <input type="text" placeholder="Имя пользователя" className="input" />
        <input type="password" placeholder="Пароль" className="input" />
        <button type="submit" className="btn">
          Войти
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708"/>
          </svg>
        </button>
      </form>
    </div>
  )
}

export default Login
