import { useEffect } from "react";
import { IMicroserviceData } from "../../types/microservice-data";
import { MicroserviceCard } from "./MicroserviceCard";
interface IProps{
    items: IMicroserviceData[];
}
export function MicroserviceCards(props: IProps) {
    return (
        <div className="m-2">
        {props.items.map((item: IMicroserviceData, index: number) => (
            <div className="mb-4 mx-4" key={index}>
                <MicroserviceCard item={item} />
            </div>
        ))}
    </div>
    );
}

