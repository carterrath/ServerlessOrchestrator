import { useEffect } from "react";
interface IProps{
    items: string[];
}
export function ScrollableList(props: IProps) {
    return (
        
        // This is your home page layout with login options
        <header className="App-header">
        <h1>Serverless Orchestrator</h1>
        <div className="login-container">
            <div style={{ maxHeight: '300px', overflowY: 'auto', border: '1px solid #ccc', padding: '5px' }}>
                <ul style={{ listStyleType: 'none', padding: 0 }}>
                    {props.items.map((item, index) => (
                        <li key={index}>{item}</li>
                    ))}
                </ul>
            </div>
        </div>
      </header>
    );
}

