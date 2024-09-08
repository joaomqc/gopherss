import { Admin, Resource } from "react-admin";
import { dataProvider } from "./dataProvider";
import {
  CategoryEdit,
  CategoryList,
  CategoryShow,
} from "./components/categories";
import { FeedEdit, FeedList, FeedShow } from "./components/feeds";
import { EntryList } from "./components/entries";
import { MyLayout } from "./components/MyLayout";
import RssFeed from "@mui/icons-material/RssFeed";
import Label from "@mui/icons-material/Label";

export const App = () => (
  <Admin title="gopherss" dataProvider={dataProvider} layout={MyLayout}>
    <Resource name="entries" list={EntryList} />
    <Resource
      name="categories"
      edit={CategoryEdit}
      list={CategoryList}
      show={CategoryShow}
      icon={Label}
    />
    <Resource
      name="feeds"
      edit={FeedEdit}
      list={FeedList}
      show={FeedShow}
      icon={RssFeed}
    />
  </Admin>
);
