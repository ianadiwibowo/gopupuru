class CreateCats < ActiveRecord::Migration[5.2]
  def change
    create_table :cats, id: false, primary_key: :id do |t|
      t.primary_key      :id, :unsigned_bigint, auto_increment: true
      t.string           :name, null: false
      t.unsigned_integer :color, limit: 1, null: false
      t.boolean          :naughty, default: false
      t.unsigned_integer :dexterity, limit: 1, null: false
      t.timestamps       null: false
    end
  end
end
