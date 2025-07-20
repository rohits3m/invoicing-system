import type { JSX } from "react";

interface IDrawerTileArgs {
    title: string;
    value: string;
    currentValue: string;
    onClick: (value: string) => void;
    icon: JSX.Element;
}

export default function DrawerTile(args: IDrawerTileArgs) {
    return(
        <div onClick={ () => args.onClick(args.value) } className={`px-6 py-2 cursor-pointer w-full hover:bg-blue-200 transition flex items-center gap-2 ${args.value == args.currentValue ? 'bg-blue-300 hover:bg-blue-300' : 'bg-white' }`}>
            <p>{args.icon}</p>
            <h1>{args.title}</h1>
        </div>
    );
}