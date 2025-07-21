import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import Api from "@/config/api";
import type { IApiResponse, IProduct } from "@/interfaces";
import { useEffect, useState } from "react";

export default function ProductList() {

    const [products, setProducts] = useState<IProduct[]>([]);

    async function getProducts() {
        try {
            const response = await Api.get("/product");
            const json = JSON.parse(response.data) as IApiResponse;

            if(!json.success) throw json.message;

            setProducts(json.data);
        }
        catch(ex) {
            console.log(ex);
            alert(ex);
        }
    }

    useEffect(() => {
        getProducts();
    }, []);

    return(
        <div className="bg-white p-5 w-[90%] max-w-[900px] rounded-lg flex flex-col gap-5 max-h-[80%]">
            
            <h1 className="text-2xl font-bold">Products</h1>

            <Input placeholder="Search.." />

            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead className="w-[100px]">S. No.</TableHead>
                        <TableHead>Product Description</TableHead>
                        <TableHead>Price</TableHead>
                        <TableHead>Stock</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {
                        products.map(product => {
                            const index = products.findIndex(p => p.id == product.id);
                            if(index == -1) return <></>;

                            return(
                                <TableRow key={product.id}>
                                    <TableCell>{index + 1}</TableCell>
                                    <TableCell>{product.title}</TableCell>
                                    <TableCell>₹ {product.price}</TableCell>
                                    <TableCell>{product.stock}</TableCell>
                                </TableRow>
                            );
                        })
                    }
                </TableBody>
            </Table>

        </div>
    );
}