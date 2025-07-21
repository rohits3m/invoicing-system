export interface IApiResponse {
    success: boolean;
    data: any;
    message?: string;
}

export interface IProduct {
    id: number;
    title: string;
    price: number;
    stock: number;
    created_on: Date;
    updated_on: Date;
}