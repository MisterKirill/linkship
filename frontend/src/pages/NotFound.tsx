import { Link } from "react-router-dom"

function NotFound() {
  return (
    <>
      <h1>404</h1>
      
      <span>Страница не найдена</span>

      <Link to="/" className="link">Вернуться домой</Link>
    </>
  )
}

export default NotFound
