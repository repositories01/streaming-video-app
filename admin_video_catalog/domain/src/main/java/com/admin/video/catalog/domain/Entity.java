package com.admin.video.catalog.domain;

import java.util.Objects;

public abstract class Entity<ID extends Identifier> {
    protected final ID id;

    public Entity(ID id) {
        Objects.requireNonNull(id, "Id should not be null");
        this.id = id;
    }

    @Override
    public boolean equals(final Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        final Entity<?> entity = (Entity<?>) o;
        return id.equals(entity.id);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
