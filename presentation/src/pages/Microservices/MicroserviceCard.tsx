import { useEffect } from "react";
import { IMicroservice } from "../../types/microservice";
interface IProps{
    item: IMicroservice;
}
export function MicroserviceCard(props: IProps) {
    return (
        <div className="p-2 bg-gray-200 rounded-xl drop-shadow-lg hover:scale-[101%] transition duration-150 ease-in-out">
            <h3>{props.item.Name}</h3>
            <p>Author: {props.item.Author}</p>
            <p>RepoLink: {props.item.RepoLink}</p>
            <p>Input: {props.item.Input}</p>
            <p>ID: {props.item.ID}</p>
            <p>CreatedAt: {props.item.CreatedAt.toString()}</p>
            <p>UpdatedAt: {props.item.UpdatedAt.toString()}</p>
            <p>DeletedAt: {props.item.DeletedAt?.toString()}</p>
            <h4>Inputs</h4>
            <ul>
                {props.item.Inputs.map((input, index) => (
                    <li key={index}>
                        <p>MicroserviceID: {input.MicroserviceID}</p>
                        <p>Id: {input.Id}</p>
                        <p>Name: {input.Name}</p>
                        <p>DataType: {input.DataType}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
}

