import { useEffect } from "react";
import { IMicroservice } from "../../types/microservice";
interface IProps{
    items: IMicroservice[];
}
export function MicroserviceCards(props: IProps) {
    return (
        <div className="scrollable-list">
        {/* Map over props.items to generate list items */}
        {props.items.map((item: IMicroservice, index: number) => (
            <div className="scrollable-list-item" key={index}>
                {}
            </div>
        ))}
    </div>
    );
}

