import { useEffect } from "react";
import { IMicroservice } from "../../types/microservice";
import { MicroserviceCard } from "./MicroserviceCard";
interface IProps{
    items: IMicroservice[];
}
export function MicroserviceCards(props: IProps) {
    return (
        <div className="m-2">
        {/* Map over props.items to generate list items */}
        {props.items.map((item: IMicroservice, index: number) => (
            <div className="mb-4 mx-4" key={index}>
                <MicroserviceCard item={item} />
            </div>
        ))}
    </div>
    );
}

