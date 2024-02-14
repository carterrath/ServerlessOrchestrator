import { useEffect } from "react";
interface IProps{
    items: string[];
}
export function ScrollableList(props: IProps) {
    return (
        
        <div className="scrollable-list">
        {/* Map over props.items to generate list items */}
        {props.items.map((item, index) => (
            <div className="scrollable-list-item" key={index}>{item}</div>
        ))}
    </div>
    );
}

