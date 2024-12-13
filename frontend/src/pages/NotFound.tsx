import { Link } from "react-router-dom"

function NotFound() {
  return (
    <>
      <h1>404</h1>
      
      <span>This page was not found</span>

      <Link to="/" className="text-link">Return home</Link>
    </>
  )
}

export default NotFound
