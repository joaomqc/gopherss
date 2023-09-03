import { Edit, SimpleForm, TextInput } from "react-admin";

export const FeedEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput source="id" />
      <TextInput source="name" />
      <TextInput source="url" />
      <TextInput source="category.id" />
    </SimpleForm>
  </Edit>
);
