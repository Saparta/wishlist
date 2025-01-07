// src/mocks/mockData.ts

export interface WishlistItem {
  id: string;
  name: string;
  url: string;
  price: number;
  is_gifted: boolean;
  gifted_by: string;
}

export interface Wishlist {
  id: number;
  title: string;
  items: WishlistItem[];
  description: string;
}

export interface SharedWishlistItem {
  [key: string]: "gotten" | "not gotten";
}

export interface SharedWishlist {
  id: number;
  title: string;
  items: WishlistItem;
  description: string;
}

export interface Friend {
  id: number;
  name: string;
  lastGifted: string;
}

export const mockWishlists: Wishlist[] = [
  {
    id: 1,
    title: "Birthday Wishlist",
    items: [
      {
        id: "1",
        name: "Watch",
        url: "http://example.com/watch",
        price: 199.99,
        is_gifted: false,
        gifted_by: "",
      },
      {
        id: "2",
        name: "Headphones",
        url: "http://example.com/headphones",
        price: 99.99,
        is_gifted: true,
        gifted_by: "John",
      },
    ],
    description: "Wishlist for my birthday",
  },
  {
    id: 2,
    title: "Holiday Wishlist",
    items: [
      {
        id: "1",
        name: "Camera",
        url: "http://example.com/camera",
        price: 499.99,
        is_gifted: true,
        gifted_by: "Jane",
      },
      {
        id: "2",
        name: "Travel Bag",
        url: "http://example.com/travel-bag",
        price: 79.99,
        is_gifted: true,
        gifted_by: "Doe",
      },
      {
        id: "3",
        name: "Shoes",
        url: "http://example.com/shoes",
        price: 59.99,
        is_gifted: false,
        gifted_by: "",
      },
    ],
    description: "Wishlist for the holiday season",
  },
  {
    id: 3,
    title: "Graduation Wishlist",
    items: [
      {
        id: "1",
        name: "Laptop",
        url: "http://example.com/laptop",
        price: 999.99,
        is_gifted: true,
        gifted_by: "Smith",
      },
      {
        id: "2",
        name: "Tablet",
        url: "http://example.com/tablet",
        price: 299.99,
        is_gifted: false,
        gifted_by: "",
      },
      {
        id: "3",
        name: "Backpack",
        url: "http://example.com/backpack",
        price: 49.99,
        is_gifted: false,
        gifted_by: "",
      },
    ],
    description: "Wishlist for graduation",
  },
];

export const mockSharedWishlists: Wishlist[] = [
  {
    id: 1,
    title: "Grandparents Wishlist",
    items: [
      {
        id: "1",
        name: "Watch",
        url: "http://example.com/watch",
        price: 199.99,
        is_gifted: true,
        gifted_by: "",
      },
      {
        id: "2",
        name: "Headphones",
        url: "http://example.com/headphones",
        price: 99.99,
        is_gifted: true,
        gifted_by: "John",
      },
    ],
    description: "Wishlist for grandparents",
  },
  {
    id: 2,
    title: "Holiday Wishlist w/kids",
    items: [
      {
        id: "1",
        name: "Camera",
        url: "http://example.com/camera",
        price: 499.99,
        is_gifted: true,
        gifted_by: "Jane",
      },
      {
        id: "2",
        name: "Travel Bag",
        url: "http://example.com/travel-bag",
        price: 79.99,
        is_gifted: true,
        gifted_by: "Doe",
      },
      {
        id: "3",
        name: "Shoes",
        url: "http://example.com/shoes",
        price: 59.99,
        is_gifted: false,
        gifted_by: "",
      },
    ],
    description: "Wishlist for the holiday season with kids",
  },
  {
    id: 3,
    title: "Twins Graduation Wishlist",
    items: [
      {
        id: "1",
        name: "Laptop",
        url: "http://example.com/laptop",
        price: 999.99,
        is_gifted: true,
        gifted_by: "Smith",
      },
      {
        id: "2",
        name: "Tablet",
        url: "http://example.com/tablet",
        price: 299.99,
        is_gifted: false,
        gifted_by: "",
      },
      {
        id: "3",
        name: "Backpack",
        url: "http://example.com/backpack",
        price: 49.99,
        is_gifted: false,
        gifted_by: "",
      },
    ],
    description: "Wishlist for twins graduation",
  },
];

export const mockFriends: Friend[] = [
  { id: 1, name: "Alice", lastGifted: "2024-12-10" },
  { id: 2, name: "Bob", lastGifted: "2024-11-05" },
];
