import DrawerTile from "@/components/drawer_tile";
import { useState } from "react";
import { FaList, FaPlus } from "react-icons/fa6";
import ProductList from "./views/product/ProductList";
import ProductCreate from "./views/product/ProductCreate";
import InvoiceList from "./views/invoice/InvoiceList";
import InvoiceCreate from "./views/invoice/InvoiceCreate";

export default function HomeScreen() {

    const views: any = {
        "product-list": <ProductList />,
        "product-create": <ProductCreate />,
        "invoice-list": <InvoiceList />,
        "invoice-create": <InvoiceCreate />,
    };

    const [selectedView, setSelectedView] = useState<string>("");

    function onDashboardTileClicked(value: string) {
        setSelectedView(value);
    }

    return(
        <div className="w-dvw h-dvw flex">

            {/* Main Container */}
            <div className="h-dvh flex-4 bg-zinc-200 flex justify-center items-center">
                { views[selectedView] }
            </div>

            {/* Sidebar */}
            <div className="h-full flex-1 bg-white">
                <div className="flex justify-center px-5 py-10 bg-blue-400 mb-3 flex-col items-center">
                    <h1 className="text-3xl font-extralight text-white mb-4">INVOICING</h1>
                    <p className="text-sm text-white">Created by:</p>
                    <p className="text-base text-white">ROHIT SEMRIWAL</p>
                </div>

                <div className="mb-2">
                    <h3 className="px-4 font-bold text-base">Product</h3>
                    <DrawerTile 
                        onClick={onDashboardTileClicked} 
                        currentValue={selectedView} 
                        value="product-list" 
                        title="Product List" 
                        icon={<FaList />} />
                        
                    <DrawerTile onClick={onDashboardTileClicked} currentValue={selectedView} value="product-create" title="Add Product" icon={<FaPlus />} />
                </div>

                <div>
                    <h3 className="px-4 font-bold text-base">Invoice</h3>
                    <DrawerTile onClick={onDashboardTileClicked} currentValue={selectedView} value="invoice-list" title="Invoice List" icon={<FaList />} />
                    <DrawerTile onClick={onDashboardTileClicked} currentValue={selectedView} value="invoice-create" title="Create Invoice" icon={<FaPlus />} />
                </div>
            </div>
            
        </div>
    );
}