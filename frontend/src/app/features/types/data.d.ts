export interface FormValues {
    name: string;
    email: string;
    password: string;
}

export interface LinkProps {
    href: string;
    children: string;
}

export interface ButtonProps {
    type: ButtonHTMLAttributes
    children: ReactNode;
}