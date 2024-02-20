export interface IMicroservice {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: Date | null;

    // Microservice specific fields
    Name: string;
    PlaceHolder: string;
    Input: string;
    Author: string;    
  }