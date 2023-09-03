import { Datagrid, List, TextField, UrlField } from "react-admin";

export const FeedList = () => (
  <List>
    <Datagrid rowClick="show">
      <TextField source="id" />
      <TextField source="name" />
      <UrlField source="url" />
      <TextField source="category.name" />
    </Datagrid>
  </List>
);
