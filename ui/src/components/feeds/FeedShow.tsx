import {
  ArrayField,
  BooleanField,
  Datagrid,
  Show,
  SimpleShowLayout,
  TextField,
  UrlField,
  useRecordContext,
} from "react-admin";

const FeedTitle = () => {
  const record = useRecordContext();
  if (record == null) {
    return null;
  }
  return <div>Feed {record.name}</div>;
};

export const FeedShow = () => (
  <Show title={<FeedTitle />}>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="name" />
      <UrlField source="url" />
      <TextField source="category.name" />
    </SimpleShowLayout>
  </Show>
);
