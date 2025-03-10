#include <stdlib.h>
#include <stdio.h>

struct node {
  int key;
  struct node* left;
  struct node* right;
};
typedef struct node* Node;

Node transform_tree_to_linked_list(Node in)
{
  // This function will return a sorted doubley linked list where the return value is
  // the first element of the linked list.
}

void add_to_tree(Node in, int key)
{
  Node new_node;
  Node* direction;

  direction =
    key < in->key
    ? &(in->left)
    : &(in->right);

  if ( *direction )
  {
    add_to_tree(*direction, key);
  }
  else
  {
    new_node = malloc(sizeof(struct node));
    new_node->left  = 0;
    new_node->right = 0;
    new_node->key   = key;

    *direction = new_node;
  }
}

void display_linked_list(Node start)
{
  Node this = start;

  printf("starting list\n");

  if ( this->left )
    printf("Assertion error: first node left is not null\n");

  while (1)
  {
    printf("Data: %d\n", this->key);

    if ( this->right )
    {
      if ( this->right->left != this )
        printf("Assertion error: reverse linking is not correct\n");

      this = this->right;
    }
    else
    {
      break;
    }
  }

  printf("ending list\n");
}

int main(int argc, char **argv)
{
  Node node_tree;
  Node node_first;

  int ii;

  int data[19] = {
    11,
    5,
    2,
    12,
    14,
    19,
    4,
    17,
    16,
    3,
    9,
    7,
    15,
    6,
    13,
    1,
    8,
    18,
    20
  };

  node_tree = malloc(sizeof(struct node));
  node_tree->left = 0;
  node_tree->right = 0;
  node_tree->key = 10;

  for ( ii = 0 ; ii < 19 ; ii++ )
    add_to_tree(node_tree, data[ii]);

  node_first = transform_tree_to_linked_list(node_tree);

  display_linked_list(node_first);
}

