
export function UpperNav(){
    return(
        <ul className="c-nav-links c-nav-left">
            <li>
                <a className="navbar-brand" href="/Home">
                    <img src="/src/assets/images/logo.png" alt="Logo" 
                    style={{maxHeight:  "35px", marginLeft: "25px", marginTop: "-8px"}}
                    /> <b> Serverless Orchestrator</b>
                </a>
            </li>
            <li>
                <a href="/Home">Home</a>
            </li>
            <li>
                <a href="/Microservices">Microservices</a>
            </li>
            <li>
                <a href="/Developer">Upload Microservice</a>
            </li>
            <li>
                <a href="/Login">Login</a>
            </li>
        </ul>
    );
}