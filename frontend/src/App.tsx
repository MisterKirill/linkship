import { BrowserRouter, Link, Routes } from 'react-router-dom'
import './App.css'

function App() {
  return (
    <BrowserRouter>
      <nav>
        <Link href="/" className="logo">Linkship</Link>

        <Link href="/login" className="nav-link">Log In</Link>
      </nav>
      <main>
        <Routes>

        </Routes>
      </main>
    </BrowserRouter>
  )
}

export default App
