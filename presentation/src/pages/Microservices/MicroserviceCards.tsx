import { useEffect } from "react";
import { IMicroserviceData } from "../../types/microservice-data";
import { MicroserviceCard } from "./MicroserviceCard";
interface IProps{
    items: IMicroserviceData[];
}
export function MicroserviceCards(props: IProps) {
    return (
        <div className="m-2">
        {/* Map over props.items to generate list items */}
        {props.items.map((item: IMicroserviceData, index: number) => (
            <div className="mb-4 mx-4" key={index}>
                <MicroserviceCard item={item} />
            </div>
        ))}
    </div>
    );
}

