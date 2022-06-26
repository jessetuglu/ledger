import { LedgerForm } from "../components/LedgerForm"
import { User } from "../types"

interface UserProps {
    user: User|undefined,
}

export const Auth:React.FC<UserProps> = ({user}) => {
    return (
        <div>
            <LedgerForm user={user}/>
        </div>
    )
}