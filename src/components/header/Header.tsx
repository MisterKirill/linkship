import styles from './Header.module.scss'
import Link from 'next/link'
import Image from 'next/image'

export default function Header() {
    return (
        <header className={styles.wrapper}>
            <Link href="/"><Image src="/linkship.svg" className={styles.logo} width={0} height={0} alt="Linkship" priority /></Link>
            <div className={styles.userspace}>
                <Link href="/register"><button>Register</button></Link>
                <Link href="/login"><button>Login</button></Link>
            </div>
        </header>
    )
}