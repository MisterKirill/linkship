import { Outlet } from "react-router-dom"
import classes from "./style.module.css"

function ProfileLayout() {
  return (
    <div className={classes.wrapper}>
      <Outlet />
    </div>
  )
}

export default ProfileLayout
