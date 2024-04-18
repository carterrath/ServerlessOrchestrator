import { IMicroserviceData } from '../../types/microservice-data';
import { MicroserviceCard } from './MicroserviceCard';
import { useState, useEffect } from 'react';

interface IProps {
  items: IMicroserviceData[];
  search: string;
  getMicroservices: () => void;
}

export function MicroserviceCards(props: IProps) {
  const [searchedItems, setSearchedItems] = useState<IMicroserviceData[]>([]);

  // Update searchedItems when items or search change
  useEffect(() => {
    // If search is empty, set searchedItems to props.items
    if (props.search.trim() === '') {
      setSearchedItems(props.items);
    } else {
      // Otherwise, filter items based on search criteria
      const filtered = props.items.filter((item) => {
        // You can modify this condition based on your search requirements
        return item.FriendlyName.toLowerCase().includes(props.search.toLowerCase());
      });
      setSearchedItems(filtered);
    }
  }, [props.items, props.search]);

  return (
    <div className="m-2">
      {searchedItems
        .sort(
          (a: IMicroserviceData, b: IMicroserviceData) =>
            new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime(),
        )
        .map((item: IMicroserviceData, index: number) => (
          <div className="mb-4 mx-4" key={index}>
            <MicroserviceCard item={item} getMicroservices={props.getMicroservices} />
          </div>
        ))}
    </div>
  );
}
