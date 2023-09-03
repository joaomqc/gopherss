import { Admin, Resource } from "react-admin";
import { dataProvider } from "./dataProvider";
import {
  CategoryEdit,
  CategoryList,
  CategoryShow,
} from "./components/category";
import { FeedEdit, FeedList, FeedShow } from "./components/feeds";
import { ArticleList } from "./components/article";
import { MyLayout } from "./components/MyLayout";
import RssFeed from "@mui/icons-material/RssFeed";
import Label from "@mui/icons-material/Label";

export const App = () => (
  <Admin title="gopherss" dataProvider={dataProvider} layout={MyLayout}>
    <Resource name="articles" list={ArticleList} />
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
