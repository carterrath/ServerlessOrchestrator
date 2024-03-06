import { useEffect } from "react";
import { IMicroserviceUpload } from "../../types/microservice-data";
import { MicroserviceCard } from "./MicroserviceCard";
interface IProps{
    items: IMicroserviceUpload[];
}
export function MicroserviceCards(props: IProps) {
    return (
        <div className="m-2">
        {/* Map over props.items to generate list items */}
        {props.items.map((item: IMicroserviceUpload, index: number) => (
            <div className="mb-4 mx-4" key={index}>
                <MicroserviceCard item={item} />
            </div>
        ))}
    </div>
    );
}

