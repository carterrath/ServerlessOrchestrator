export interface IMicroservice {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: Date | null;

    // Microservice specific fields
    Name: string;
    RepoLink: string;
    Input: string;
    Author: string; 
    Inputs   
  }