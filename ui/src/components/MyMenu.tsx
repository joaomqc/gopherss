import { Link, Menu, localStorageStore } from "react-admin";
import Divider from "@mui/material/Divider";
import SubMenu from "./SubMenu";
import { dataProvider } from "../dataProvider";
import { useEffect, useState } from "react";
import More from "@mui/icons-material/More";
import { useLocation } from "react-router-dom";

export const MyMenu = () => {
  const location = useLocation();
  const sidebarOpen = localStorageStore().getItem("sidebar.open");
  const [categories, setCategories] = useState<any[]>([]);
  useEffect(() => {
    async function getCategories() {
      const res = await dataProvider.getList("categories", {
        filter: null,
        pagination: { page: 0, perPage: 100 },
        sort: { field: "name", order: "asc" },
      });

      setCategories(res.data);
    }

    getCategories();
  }, []);

  return (
    <Menu>
      <Menu.ResourceItem name="articles" />
      <Menu.ResourceItem name="categories" />
      <Menu.ResourceItem name="feeds" />
      <Divider />
      {sidebarOpen &&
        categories.map((category: any) => (
          <SubMenu
            selected={location.search == `?category=${category.name}`}
            key={category.id}
            primaryText={category.name}
            to={`/articles?category=${category.name}`}
          >
            {category.feeds.map((feed: any) => (
              <Menu.Item
                selected={location.search == `?feed=${feed.name}`}
                key={feed.id}
                to={`/articles?feed=${feed.name}`}
                primaryText={feed.name}
              />
            ))}
          </SubMenu>
        ))}
      {!sidebarOpen && (
        <Menu.Item
          style={{ pointerEvents: "none" }}
          to=""
          leftIcon={<More />}
        />
      )}
    </Menu>
  );
};
