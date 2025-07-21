import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Api from "@/config/api";
import type { IApiResponse, IProduct } from "@/interfaces";
import { useState } from "react";

export default function ProductCreate() {

    const [form, setForm] = useState<Partial<IProduct>>({
        title: "",
        price: 0,
        stock: 0
    });

    async function submit(e: React.FormEvent) {
        e.preventDefault();
        
        try {
            const response = await Api.post("/product", JSON.stringify(form));
            const json = JSON.parse(response.data) as IApiResponse;

            if(!json.success) throw json.message;

            alert(`CREATED PRODUCT: ${json.data}`);
        }
        catch(ex) {
            console.log(ex);
            alert(ex);
        }
    }
    
    return(
        <form onSubmit={submit} className="w-[90%] max-w-[700px] bg-white p-5 flex flex-col gap-5 rounded-lg">
            <h1 className="text-2xl font-bold">Create Product</h1>
            
            <div>
                <Label>Title</Label>
                <Input onChange={e => {
                    setForm((prev) => ({
                        ...prev,
                        title: e.target.value
                    }));
                }} placeholder="Title" required autoFocus />
            </div>

            <div className="flex gap-5">
                <div className="flex-1">
                    <Label>Price</Label>
                    <Input onChange={e => {
                    setForm((prev) => ({
                        ...prev,
                        price: parseFloat(e.target.value) || 0
                    }));
                }} placeholder="Price" required />
                </div>

                <div className="flex-1">
                    <Label>Stock</Label>
                    <Input onChange={e => {
                    setForm((prev) => ({
                        ...prev,
                        stock: parseInt(e.target.value) || 0
                    }));
                }} placeholder="Stock" required />
                </div>
            </div>

            <Button>Save</Button>
        </form>
    );
}