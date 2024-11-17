import { Link } from 'react-router-dom'
import LinkButton from '../../components/LinkButton'
import './style.css'

function Profile() {
  return (
    <>
      <img src="https://github.com/MisterKirill.png" alt="MisterKirill's avatar" className="avatar" />

      <h1 className="display-name">Kirill Siukhin</h1>

      <span className="bio">
        Hi! 👋<br />
        My name is Kirill, and I'm a cool programmer from Russia.<br /><br />

        I make fun things, such as this website :).<br />
        Linkship is a very simple website where you can easily create and share your profile.<br /><br />

        Btw, these are my socials where you can reach me:
      </span>
      
      <div className="link-buttons-list">
        <LinkButton backgroundColor="229ED9" href="https://t.me/misterkirill1">Telegram</LinkButton>
        <LinkButton backgroundColor="000000" href="https://github.com/MisterKirill">GitHub</LinkButton>
      </div>

      <Link to="/" className="credit-link">Powered by <b>Linkship</b></Link>
    </>
  )
}

export default Profile
